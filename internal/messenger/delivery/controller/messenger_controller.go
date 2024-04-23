package controller

import (
	"encoding/json"
	"github.com/Point-AI/backend/config"
	"github.com/Point-AI/backend/internal/messenger/delivery/model"
	_interface "github.com/Point-AI/backend/internal/messenger/domain/interface"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type MessengerController struct {
	messengerService _interface.MessengerService
	websocketService _interface.WebsocketService
	config           *config.Config
}

func NewMessengerController(cfg *config.Config, messengerService _interface.MessengerService, websocketService _interface.WebsocketService) *MessengerController {
	return &MessengerController{
		messengerService: messengerService,
		config:           cfg,
	}
}

func (mc *MessengerController) RegisterBotIntegration(c echo.Context) error {
	userId := c.Request().Context().Value("userId").(primitive.ObjectID)
	var request model.RegisterBotRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	if err := mc.messengerService.RegisterBotIntegration(userId, request.BotToken, request.WorkspaceId); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, model.SuccessResponse{Message: "bot added successfully"})
}

func (mc *MessengerController) HandleBotMessage(c echo.Context) error {
	token := c.Param("token")
	var update *tgbotapi.Update
	if err := json.NewDecoder(c.Request().Body).Decode(&update); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	if err := mc.messengerService.HandleTelegramBotMessage(token, update); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return nil
}

func (mc *MessengerController) WSHandler(c echo.Context) error {
	userId := c.Request().Context().Value("userId").(primitive.ObjectID)
	workspaceId := c.Param("id")

	err := mc.messengerService.ValidateUserInWorkspace(userId, workspaceId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	ws, err := mc.websocketService.UpgradeConnection(c.Response(), c.Request(), workspaceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	defer mc.websocketService.RemoveConnection(workspaceId, ws)

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		var receivedMessage model.MessageRequest
		if err := json.Unmarshal(message, &receivedMessage); err != nil {
			continue
		}

		if receivedMessage.Source == "telegramBot" {
			if err := mc.messengerService.HandleTelegramPlatformMessage(userId, workspaceId, receivedMessage); err != nil {
				continue
			}

		}
	}

	return nil
}

func (mc *MessengerController) ReassignTicketToMember(c echo.Context) error {
	ticketId, workspaceId, userEmail := c.Param("ticket_id"), c.Param("id"), c.Param("email")
	userId := c.Request().Context().Value("userId").(primitive.ObjectID)

	if err := mc.messengerService.ReassignTicketToMember(userId, ticketId, workspaceId, userEmail); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}
}
