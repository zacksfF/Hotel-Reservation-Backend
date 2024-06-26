// package api

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/zacksfF/Hotel-Reservation-Backend/db/fixture"
// 	"github.com/zacksfF/Hotel-Reservation-Backend/types"
// )

// func TestUserGetBooking(t *testing.T) {
// 	db := setup(t)
// 	defer db.teardown(t)

// 	var (
// 		nonAuthUser    = fixture.AddUser(db.store, "Jimmy", "watercooler", false)
// 		user           = fixture.AddUser(db.store, "james", "foo", false)
// 		hotel          = fixture.AddHotel(db.store, "bar hotel", "a", 4, nil)
// 		room           = fixture.AddRoom(db.store, "small", true, 4.4, hotel.ID)
// 		from           = time.Now()
// 		till           = from.AddDate(0, 0, 5)
// 		booking        = fixture.AddBooking(db.Store, user.ID, room.ID, from, till)
// 		app            = fiber.New(fiber.Config{ErrorHandler: fiber.DefaultErrorHandler})
// 		route          = app.Group("/", JWTAuthentication(db.UserStore))
// 		bookingHandler = NewBookingHandler(db.store)
// 	)
// 	route.Get("/:id", bookingHandler.HandleGetBooking)
// 	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
// 	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		t.Fatalf("non 200 code got %d", resp.StatusCode)
// 	}
// 	var bookingResp *types.Booking
// 	if err := json.NewDecoder(resp.Body).Decode(&bookingResp); err != nil {
// 		t.Fatal(err)
// 	}
// 	if bookingResp.ID != booking.ID {
// 		t.Fatalf("expected %s got %s", booking.ID, bookingResp.ID)
// 	}
// 	if bookingResp.UserID != booking.UserID {
// 		t.Fatalf("expected %s got %s", booking.UserID, bookingResp.UserID)
// 	}
// 	req = httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
// 	req.Header.Add("X-Api-Token", CreateTokenFromUser(nonAuthUser))
// 	resp, err = app.Test(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if resp.StatusCode == http.StatusOK {
// 		t.Fatalf("expected a non 200 status code got %d", resp.StatusCode)
// 	}
// }

// func TestAdminGetBookings(t *testing.T) {
// 	db := setup(t)
// 	defer db.teardown(t)

// 	var (
// 		adminUser      = fixture.AddUser(db.store, "admin", "admin", true)
// 		user           = fixture.AddUser(db.Store, "james", "foo", false)
// 		hotel          = fixture.AddHotel(db.Store, "bar hotel", "a", 4, nil)
// 		room           = fixture.AddRoom(db.Store, "small", true, 4.4, hotel.ID)
// 		from           = time.Now()
// 		till           = from.AddDate(0, 0, 5)
// 		booking        = fixture.AddBooking(db.Store, user.ID, room.ID, from, till)
// 		app            = fiber.New(fiber.Config{ErrorHandler: ErrorHandler})
// 		admin          = app.Group("/", JWTAuthentication(db.User), AdminAuth)
// 		bookingHandler = NewBookingHandler(db.S)
// 	)
// 	admin.Get("/", bookingHandler.HandleGetBookings)
// 	req := httptest.NewRequest("GET", "/", nil)
// 	req.Header.Add("X-Api-Token", CreateTokenFromUser(adminUser))
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		t.Fatalf("non 200 response got %d", resp.StatusCode)
// 	}
// 	var bookings []*types.Booking
// 	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
// 		t.Fatal(err)
// 	}
// 	if len(bookings) != 1 {
// 		t.Fatalf("expected 1 booking got %d", len(bookings))
// 	}
// 	have := bookings[0]
// 	if have.ID != booking.ID {
// 		t.Fatalf("expected %s got %s", booking.ID, have.ID)
// 	}
// 	if have.UserID != booking.UserID {
// 		t.Fatalf("expected %s got %s", booking.UserID, have.UserID)
// 	}

// 	// test non-admin cannot access the bookings
// 	req = httptest.NewRequest("GET", "/", nil)
// 	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
// 	resp, err = app.Test(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if resp.StatusCode != http.StatusUnauthorized {
// 		t.Fatalf("expected status unauthorized but got %d", resp.StatusCode)
// 	}
// }