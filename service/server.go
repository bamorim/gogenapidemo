package service

import (
	"context"
	"errors"

	"github.com/bamorim/gogenapidemo/api"
	"github.com/google/uuid"
)

type Service struct {
	widgets map[string]api.Widget
}

func NewService() (*Service, error) {
	return &Service{
		widgets: make(map[string]api.Widget),
	}, nil
}

var (
	ErrWidgetNotFound = errors.New("widget not found")
	ErrInvalidParams  = errors.New("invalid parameters")
)

func (*Service) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response: api.Error{
			Message: err.Error(),
			Code:    1,
		},
	}
}

func (s *Service) WidgetsAnalyze(ctx context.Context, params api.WidgetsAnalyzeParams) (*api.AnalyzeResult, error) {
	if params.ID == "" {
		return nil, ErrInvalidParams
	}

	widget, exists := s.widgets[params.ID]

	if !exists {
		return nil, ErrWidgetNotFound
	}

	var analysis string

	switch widget.Color {
	case "red":
		analysis = "This widget is red, indicating it is likely to be high priority."
	case "blue":
		if widget.Weight > 10 {
			analysis = "This widget is blue and heavy, suggesting it is durable."
		} else {
			analysis = "This widget is blue and lightweight, indicating it is easy to handle."
		}
	default:
		analysis = "This widget has an unknown color, further analysis is needed."
	}

	// Simulate an analysis result
	result := &api.AnalyzeResult{
		ID:       params.ID,
		Analysis: analysis,
	}

	return result, nil
}

func (s *Service) WidgetsCreate(ctx context.Context, req *api.WidgetCreateRequest) (*api.Widget, error) {
	widget := api.Widget{
		ID:     uuid.New().String(),
		Color:  api.WidgetColor(req.Color),
		Weight: req.Weight,
	}

	s.widgets[widget.ID] = widget
	return &widget, nil
}

func (s *Service) WidgetsDelete(ctx context.Context, params api.WidgetsDeleteParams) error {
	_, exists := s.widgets[params.ID]
	if !exists {
		return ErrWidgetNotFound
	}

	delete(s.widgets, params.ID)
	return nil
}

func (s *Service) WidgetsList(ctx context.Context) (*api.WidgetList, error) {
	widgetList := &api.WidgetList{
		Items: make([]api.Widget, 0, len(s.widgets)),
	}

	for _, widget := range s.widgets {
		widgetList.Items = append(widgetList.Items, widget)
	}

	return widgetList, nil
}

func (s *Service) WidgetsRead(ctx context.Context, params api.WidgetsReadParams) (*api.Widget, error) {
	widget, exists := s.widgets[params.ID]
	if !exists {
		return nil, ErrWidgetNotFound
	}

	return &widget, nil
}

func (s *Service) WidgetsUpdate(ctx context.Context, req *api.WidgetUpdateRequest, params api.WidgetsUpdateParams) (*api.Widget, error) {
	widget, exists := s.widgets[params.ID]
	if !exists {
		return nil, ErrWidgetNotFound
	}

	if req.Color.IsSet() {
		widget.Color = api.WidgetColor(req.Color.Value)
	}

	if req.Weight.IsSet() {
		widget.Weight = req.Weight.Value
	}

	s.widgets[params.ID] = widget
	return &widget, nil
}

var _ api.Handler = (*Service)(nil)
