# Instagram Scraper with Real-Time Hashtag Tracking

This project is a simple web scraper that extracts posts from Instagram in real-time based on a specified hashtag. The scraper retrieves posts containing the hashtag and stores relevant information, such as text, code, and timestamp, for further processing or analysis.

## Features

- **Real-time scraping**: Continuously monitors Instagram for new posts related to a specified hashtag.
- **Hashtag-based filtering**: Only fetch posts that include the requested hashtag.
- **Post data extraction**: Extracts the post text, associated code, and timestamp.
- **Efficient data storage**: Stores posts in a structured format (e.g., JSON) for further use.
- **Error handling**: Handles potential network or rate limit errors.

## Technologies Used

- **Go (Golang)**: Backend programming language.
- **Instagram Web Scraping**: Uses Instagram’s web page structure to fetch posts.
- **JSON**: Data format for storing and transferring post details.
- **GORM**: ORM used for database interaction if needed for storing data.

## Installation

1. Clone the repository:
   ```bash   git clone https://github.com/braquetes/instagram-scraper.git
   cd instagram-scraper
