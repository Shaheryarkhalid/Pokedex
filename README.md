# 🧢 Pokédex CLI

A terminal-based Pokédex game written in Go!  
Explore the Pokémon world, catch and inspect Pokémon, and build your own Pokédex — all through an interactive command-line interface.

---

## 🚀 Features

- 🌍 View and navigate Pokémon location areas
- 🎯 Explore specific locations for available Pokémon
- 🎒 Catch and store Pokémon in your personal Pokédex
- 📊 Inspect caught Pokémon stats (type, height, weight, etc.)
- ⚡ Fast access via in-memory caching
- 💬 Helpful and interactive CLI command system

---

## 📦 Project Structure

Pokedex/

    ├── cmd/ # CLI entry point (main.go)
    
    ├── internal/

        ├── pokecache/ # In-memory TTL-based cache
        │── pokemon/ # API client and user Pokédex logic



---

## 🛠️ Getting Started

### 🔧 Prerequisites

- Go 1.20+ installed: [https://golang.org/dl/](https://golang.org/dl/)

### 📦 Build & Run

```bash
git clone git@github.com:Shaheryarkhalid/Pokedex.git
cd Pokedex
go build -o pokedex
./pokedex
```

## 🕹️ CLI Commands

| Command   | Description                                      |
| --------- | ------------------------------------------------ |
| `map`     | View the next 20 Pokémon location areas          |
| `mapb`    | Go back to the previous 20 locations             |
| `explore` | Explore a specific location to see local Pokémon |
| `catch`   | Try to catch a Pokémon by name                   |
| `inspect` | View details of a Pokémon you’ve caught          |
| `pokedex` | List all caught Pokémon                          |
| `help`    | Show all available commands                      |
| `clear`   | Clear the terminal screen                        |
| `exit`    | Quit the Pokédex game                            |


## 🧠 How It Works
Uses the PokéAPI to fetch live Pokémon and location data.

Caches API results in memory for faster repeated lookups.

Stores caught Pokémon locally during your session.

Processes and validates user input with a clean CLI command system.

## 🧪 Example Usage


Pokedex > map
Pokedex > explore viridian-forest
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex

## 🙌 Credits
Pokémon data provided by PokéAPI

Built with Go for educational and entertainment purposes

## 🛡️ License
This project is licensed under the MIT License.

