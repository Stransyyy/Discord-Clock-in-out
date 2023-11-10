# Clocking-in
 
This program will help me in my working environment, it will check the time I get on my computer at work (in the future it will run automatically as soon as I open Discord... that's the idea) and the time I stop using it and get ready to leave the office

## Contents

- [Clocking-in](#clocking-in)
- [Description](#description)
- [Overview](#Overview)
- [Features](#features)
- [Benefits](#benefits)
- [Usage](#Usage)

## Description

# Automatic Clock-In/Out Discord Bot with Database Integration

## Overview

This Go program simplifies the process of tracking work hours by combining the power of a database and a Discord bot. It enables employees to easily clock in and out of their work shifts, while automatically notifying their supervisor on a Discord channel about their work hours.

## Features

- **Database Integration**: Utilizes a database (e.g., MySQL, SQLite) to store clock-in and clock-out records for employees, including their unique identifier, name, clock-in time, and clock-out time.

- **Discord Bot**: Incorporates a Discord bot that allows users to interact with it through text commands. Employees can use commands like `!clock-in` and `!clock-out` to manage their work shifts.

- **Automatic Notifications**: After an employee clocks in or out, the bot automatically sends a notification message to a designated Discord channel, including the employee's name, action (clock-in or clock-out), and timestamp.

- **Admin Features**: Provides administrative capabilities for managing employee records and attendance data. Admins can view employee attendance history, add new employees, and update existing records.

- **Security and Data Integrity**: Ensures data security with user authentication for admin functions and maintains data integrity through validation and error handling mechanisms.

## Benefits

- Simplifies work-hour tracking for employees.
- Enhances transparency in attendance management.
- Reduces manual reporting efforts.
- Provides a user-friendly Discord interface for employees.
- Automates notifications to keep supervisors informed.
- Stores attendance records in a reliable database for reference.
- Adds up all the worked hours in the worked period

## Usage

Will have commands for the user and admin end of the application. Like `!clock-in`, `!clock-out`, `!clocked-out-at <time clocked out> ` <- this if for some reason the employee forgets to clock out from work, `!hours-worked <user> <from(date)> <to(date)>`, etc.
