# WanderFareTracker

**Find the best travel deals to make your dream journey affordable.**

---

## Project Overview

WanderFareTracker is an application designed to help users track flight fares and hotel discounts, aiming to assist individuals in realizing their travel dreams at lower prices. Whether you're a backpacker or a family traveler, this tool will regularly monitor various travel websites and automatically notify users when low prices are detected.

### Motivation

This project is aimed at helping those with travel dreams but limited budgets. It allows them the opportunity to achieve their travel goals at affordable prices. It can also help people who want to travel while working, explore the world, and experience new cultures without breaking the bank.

---

## Features

- **Flight Fare Tracking**: Track and get notified when flight prices drop for specific routes (e.g., city pairs).
- **Custom Alerts**: Set up custom notifications when prices go below a specified threshold.
- **Flexible Travel Preferences**: Choose from direct flights, multi-stop flights, and exclude specific countries for layovers (e.g., avoid China for layovers).
- **Subscription with Expiry Date**: Set a date for when the subscription to the fare alert expires. If no deals are found by that date, the subscription will be automatically canceled.
- **Destination Recommendations**: If you don’t have a specific destination in mind, the app can suggest the cheapest destinations departing from your location.
- **Price Alerts via Telegram**: Receive timely alerts on price changes via Telegram.
- **Historical Price Data**: Use historical data to define "cheap" prices based on past trends.
- **Customizable Notification Threshold**: Define the price at which you'd like to be notified about flight deals.
- **Web Interface for Easy Access**: Simple interface to manage subscriptions and preferences for users.

---

## Tech Stack

- **Backend**:
  - **Golang**
  - **PostgreSQL**
  - **Redis**
  - **Kafka**
  - **Python**

- **Frontend**:
  - **React.js**

- **Infrastructure & Deployment**:
  - **Docker**
  - **Kubernetes**
  - **AWS**
---

## System Architecture

1. **Frontend** (React.js): 
   - User interacts with the web interface to set up subscription preferences.
   
2. **API Layer** (Golang):
   - Handles user inputs, such as storing subscription data, and triggers the relevant background tasks.

3. **Web Scraper** (Python):
   - Periodically scrapes data from flight websites (e.g., Google Flights, Skyscanner, etc.).
   - Processes flight data, checks for price drops or favorable deals based on user preferences.
   
4. **Database** (PostgreSQL):
   - Stores flight details, user data, and subscription details.
   - Historical price data is stored to help define what constitutes a "cheap" flight.
   
5. **Message Queue** (Kafka): 
   - Used to manage communication between the flight scraper and notification system.
   
6. **Notification System** (Telegram Bot API):
   - Once a price alert is triggered (based on user preferences), the system sends notifications to users via Telegram.

