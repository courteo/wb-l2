package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestParseGood(t *testing.T) {
	date, _ :=time.Parse("2006-01-02", "2022-01-01")
	resEvent := Event{
		ID: 0,
		UserID: 1,
		Title: "TestEvent",
		Date: date,
	}
	form := strings.NewReader("user_id=1&date=2022-01-01&title=TestEvent")
	req := httptest.NewRequest("POST", "/create_event", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ev, err :=ParseInputDataToEvent(req)
	// fmt.Println(ev, err)
	if resEvent != *ev || err != nil {
		t.Errorf("Expected event %v,\n but got %v\n Expected error %v,\n but got %v\n", resEvent, *ev, nil, err)
	}
}

func TestParseBad(t *testing.T) {
	form := strings.NewReader("user id=1&date=2022-01-01&title=TestEvent")
	req := httptest.NewRequest("POST", "/create_event", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ev, err :=ParseInputDataToEvent(req)
	// fmt.Println(ev, err)
	if nil != ev || err == nil {
		t.Errorf("Expected event %v,\n but got %v\n Expected error %v,\n but got %v\n", nil, *ev, nil, err)
	}
}

func TestCreateEventHandler(t *testing.T) {

	// массив куда записываются данные
	Calender.Events = make([]*Event, 0)
	date, _ :=time.Parse("2006-01-02", "2022-01-01")
	resEvent := Event{
		ID: 0,
		UserID: 1,
		Title: "TestEvent",
		Date: date,
	}
	// Создаем тестовый запрос
	form := strings.NewReader("user_id=1&date=2022-01-01&title=TestEvent")
	req := httptest.NewRequest("POST", "/create_event", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Создаем ResponseWriter для записи ответа
	rec := httptest.NewRecorder()

	// Вызываем обработчик
	CreateEventHandler(rec, req)
	// Проверяем статус код и содержимое ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}
	// проверяем добавился ивент
	if len(Calender.Events) !=1 || *Calender.Events[0] != resEvent {
		t.Errorf("Expected len %v,\n but got %v,\n Expected Calendar.Events %v,\n but got %v\n", 1,len(Calender.Events), resEvent,*Calender.Events[0] )
	}

	expectedResponse := fmt.Sprintf("%s\n", `{"result":"Event created"}`)
	if rec.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rec.Body.String())
	}
}

func TestUpdateEventHandler(t *testing.T) {
	// массив который надо обновить
	Calender.Events = make([]*Event, 0)
	date, _ :=time.Parse("2006-01-02", "2022-01-01")
	event := Event{
		ID: 0,
		UserID: 1,
		Title: "TestEvent",
		Date: date,
	}
	Calender.Events = append(Calender.Events, &event)
	resEvent := Event{
		ID: 0,
		UserID: 1,
		Title: "Event",
		Date: date,
	}
	// Создаем тестовый запрос
	form := strings.NewReader("user_id=1&date=2022-01-01&title=Event")
	req := httptest.NewRequest("POST", "/update_event", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Создаем ResponseWriter для записи ответа
	rec := httptest.NewRecorder()

	// Вызываем обработчик
	UpdateEventHandler(rec, req)
	// Проверяем статус код и содержимое ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}
	// проверяем добавился ивент
	if len(Calender.Events) !=1 || *Calender.Events[0] != resEvent {
		t.Errorf("Expected len %v,\n but got %v,\n Expected Calendar.Events %v,\n but got %v\n", 1,len(Calender.Events), resEvent,*Calender.Events[0] )
	}

	expectedResponse := fmt.Sprintf("%s\n", `{"result":"Event updated"}`)
	if rec.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rec.Body.String())
	}
}

func TestDeleteEventHandler(t *testing.T) {
	// массив откуда надо удалять
	Calender.Events = make([]*Event, 0)
	date, _ :=time.Parse("2006-01-02", "2022-01-01")
	event := Event{
		ID: 0,
		UserID: 1,
		Title: "TestEvent",
		Date: date,
	}
	Calender.Events = append(Calender.Events, &event)
	// Создаем тестовый запрос
	form := strings.NewReader("user_id=1&date=2022-01-01&title=TestEvent")
	req := httptest.NewRequest("POST", "/delete_event", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Создаем ResponseWriter для записи ответа
	rec := httptest.NewRecorder()

	// Вызываем обработчик
	DeleteEventHandler(rec, req)
	// Проверяем статус код и содержимое ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}
	// проверяем добавился ивент
	if len(Calender.Events) !=0 {
		t.Errorf("Expected len %v,\n but got %v,\n", 0,len(Calender.Events))
	}

	expectedResponse := fmt.Sprintf("%s\n", `{"result":"Event deleted"}`)
	if rec.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rec.Body.String())
	}
}


func TestEventsForDayHandler(t *testing.T) {
	// Добавляем события в календарь
	event1 := &Event{
		ID:     1,
		UserID: 1,
		Date:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		Title:  "Event 1",
	}
	event2 := &Event{
		ID:     2,
		UserID: 2,
		Date:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		Title:  "Event 2",
	}
	Calender.Events = []*Event{event1, event2}

	// Создаем тестовый запрос
	req := httptest.NewRequest("GET", "/events_for_day?date=2022-01-01", nil)

	// Создаем ResponseWriter для записи ответа
	rec := httptest.NewRecorder()

	// Вызываем обработчик
	EventsForDayHandler(rec, req)

	// Проверяем статус код и содержимое ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	expectedResponse := `{"result":"events for day [{\"id\":1,\"user_id\":1,\"date\":\"2022-01-01T00:00:00Z\",\"title\":\"Event 1\"},{\"id\":2,\"user_id\":2,\"date\":\"2022-01-01T00:00:00Z\",\"title\":\"Event 2\"}]"}`
	expectedResponse = fmt.Sprintf("%s\n", expectedResponse)
	if rec.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rec.Body.String())
	}
}


func TestEventsForMonthHandler(t *testing.T) {
	// Добавляем события в календарь
	event1 := &Event{
		ID:     1,
		UserID: 1,
		Date:   time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
		Title:  "Event 1",
	}
	event2 := &Event{
		ID:     2,
		UserID: 2,
		Date:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		Title:  "Event 2",
	}
	Calender.Events = []*Event{event1, event2}

	// Создаем тестовый запрос
	req := httptest.NewRequest("GET", "/events_for_month?date=2022-01-01", nil)

	// Создаем ResponseWriter для записи ответа
	rec := httptest.NewRecorder()

	// Вызываем обработчик
	EventsForMonthHandler(rec, req)

	// Проверяем статус код и содержимое ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	expectedResponse := `{"result":"Events for month [{\"id\":2,\"user_id\":2,\"date\":\"2022-01-01T00:00:00Z\",\"title\":\"Event 2\"}]"}`
	expectedResponse = fmt.Sprintf("%s\n", expectedResponse)
	if rec.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rec.Body.String())
	}
}

func TestEventsForWeekHandler(t *testing.T) {
	// Добавляем события в календарь они на разных неделях
	event1 := &Event{
		ID:     1,
		UserID: 1,
		Date:   time.Date(2022, 1, 3, 0, 0, 0, 0, time.UTC),
		Title:  "Event 1",
	}
	event2 := &Event{
		ID:     2,
		UserID: 2,
		Date:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		Title:  "Event 2",
	}
	Calender.Events = []*Event{event1, event2}

	// Создаем тестовый запрос
	req := httptest.NewRequest("GET", "/events_for_week?date=2022-01-01", nil)

	// Создаем ResponseWriter для записи ответа
	rec := httptest.NewRecorder()

	// Вызываем обработчик
	EventsForWeekHandler(rec, req)

	// Проверяем статус код и содержимое ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	expectedResponse := `{"result":"Events for week [{\"id\":2,\"user_id\":2,\"date\":\"2022-01-01T00:00:00Z\",\"title\":\"Event 2\"}]"}`
	expectedResponse = fmt.Sprintf("%s\n", expectedResponse)
	if rec.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rec.Body.String())
	}
}