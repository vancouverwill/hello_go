package main

import (
    "errors"
    "regexp"
    "strconv"
    "strings"
)

const degreesInCircle = 360

/*
GetAngleBetweenTimeString of a string of the form 12:30 
i.e one that is a valid digital
*/
func GetAngleBetweenTimeString(time string) (int, error) {
    if !isValidTime(time) {
        return -1, errors.New("Not valid time string")
    }
    strings := strings.Split(time, ":")
    hour, hourError := strconv.Atoi(strings[0]) 
    min, minError := strconv.Atoi(strings[1]) 
    if hourError != nil || minError != nil {
        return -1, errors.New("Not valid time string")
    }
    return GetAngleBetweenClockHands(hour, min)
}

/*
GetAngleBetweenClockHands is a simple Go function to return the angle in degrees

steps 
1 - determine hour angle 
2 - determine min angle
3 - return minimum angle between clock hands
*/
func GetAngleBetweenClockHands(hour int, min int) (int, error) {
    var err error
    
    if hour < 0 || hour > 24 {
        return -1, errors.New("hour out of range")
    }
    if min < 0 || min > 60 {
        return -1, errors.New("min out of range")
    }
    
    hourAngle := getHourAngle(hour)
    minAngle := getMinuteAngle(min)
    var difference int
    if hourAngle > minAngle {
        difference = hourAngle - minAngle
    } else {
        difference =  minAngle - hourAngle
    }
   
    if difference > degreesInCircle / 2 { 
        difference = degreesInCircle - difference
    } 
    
    return difference, err
}

func getHourAngle(hour int) (angle int) {
    angle = hour * degreesInCircle / 12 
    return angle
}

func getMinuteAngle(minute int) (angle int) {
    angle = minute * degreesInCircle  / 60 
    return angle
}

func isValidTime(time string) bool {
    reg := regexp.MustCompile(`^[01]{1}[0-9]{1}|2{1}[0-4]{1}\:\d{2}$`)
	return reg.MatchString(time)
}