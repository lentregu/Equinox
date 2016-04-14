# Welcome to Human Authorization

This project aims to build your own low cost authorization tool by means of the face detection.

You can find here [github project] (https://github.com/lentregu/Equinox) the code (result of a Hackaton).


![Architecture] (img/equinox.jpg)

The picture above shows the architecture employed to implement our **_Human Authorization_** low cost tool.

+ Movement Sensor
+ Arduino
+ Raspberry Pi
+ Raspicam
+ Microsoft Oxfor Project API
+ Screen



### Arduino

This component is in charge to detect presence and communicate it to the Raspberry Server in order to be authorized.

## Movement Sensor

It is connected to the arduino, this sensor detects movement (human presence) and sends a digital signal to an Arudino's digital input PIN.

## Movement Manager

This is a little program that is listening the input PIN (0|1) where the movement sensor writes, and resend it to the Raspberry Pi through the serial port where will be captured by the PhotoManager.

## Raspberry Pi

This component acts as the server for the Human Authorization 

## PhotoManager

This program is listening in the serial port and when detects a 1, it asks the Raspicam to take a photo and store it in a file. Furthermore, it sends a request to the Human Authorization server to authorize the person whose photo is stored in the file provided in the request (query parameter).

## Human Authorization

This component is a server that exposes an API to authorize a person if him/her belongs to a list o authorized people that is stored in the Azure Cloud. The API receives a photo of the person to be authenticated and it makes a request to faceAPI of Microsoft that will return a list of possible person in the list with a percentage of matching with the person to be authenticated.

The Human Authorization server applies an algorithm to determine if the user is one of them and if so it writes "AUTHORIZED" in the screen otherwise it writes "UNAUTHORIZED" at the same time it can send an SMS to a configured number with the authorization result.

## Provision Tool

This tool is a CLI tool that allows us to creat a list of authorised people, to add a person to a list of people, to list all the exixting "people list" and to list all the people in a specific list.

## Raspicam

Camera connected to the Raspberry Pi to take photos to the people to be authorised.

## Microsoft Oxfor Project

Microsoft IA server that offers API to detect faces, create list of faces, find similar faces, etc.

## Screen

Device used to write the result of the authorization.
