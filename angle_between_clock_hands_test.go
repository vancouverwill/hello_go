package hello_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAngleBetweenTimeString(t *testing.T) {
    expected := 60
    actual, _ := GetAngleBetweenTimeString("07:45")
    
    assert.Equal(t, expected, actual, "return angle between time string like 07:45")
}

func TestMinuteAheadOfHourHand(t *testing.T) {
    expected := 60
    actual, _ := GetAngleBetweenClockHands(7, 45)
    
    assert.Equal(t, expected, actual, "Minute 45 is ahead of hours at 7 and should return 60 degrees difference")  
}

func TestHourAheadOfMinuteHand(t *testing.T) {
    expected := 120
    actual, _ := GetAngleBetweenClockHands(5, 5)
    
    assert.Equal(t, expected, actual, "Minute 5 is behind of hours at 5 and should return 120 degrees difference")  
}

func TestHourAheadOfMinuteAndCrossesMidNightHand(t *testing.T) {
    expected := 60
    actual, _ := GetAngleBetweenClockHands(11, 5)
    
    assert.Equal(t, expected, actual, "Minute 5 is behind of hours at 11 but shortest distance is across midnight at 60")  
}


func TestValidTime(t *testing.T) {
    assert.Equal(t, true, isValidTime("05:00"), "Test valid time")
    assert.Equal(t, true, isValidTime("16:00"), "Test valid time")
    assert.Equal(t, true, isValidTime("20:00"), "Test valid time 20")
    assert.Equal(t, true, isValidTime("24:00"), "Test valid time 24")
    assert.Equal(t, true, isValidTime("00:00"), "Test valid time")
}

func TestInValidTime(t *testing.T) { 
    assert.Equal(t, false, isValidTime("30:00"), "Test invalid time")
    assert.Equal(t, false, isValidTime("25:00"), "Test invalid time")
    assert.Equal(t, false, isValidTime("ee:00"), "Test invalid time")
    assert.Equal(t, false, isValidTime("-13:00"), "Test invalid time")
}