# Kana to Eisu

A keyboard input monitoring tool for macOS that automatically switches to English input mode when there is no keyboard input for 3 seconds.

## Features

- Real-time keyboard input monitoring
- Automatically switches to English input mode after 3 seconds of inactivity
- Efficient background operation

## Requirements

- macOS
- Go 1.21 or higher
- Accessibility permissions
- im-select

## Installation

1. Clone the repository:

```zsh
git clone https://github.com/h7kayama/kana_to_eisu.git
cd kana_to_eisu
```

2. Install dependencies:

```zsh
go mod tidy
```

3. Install im-select:

```zsh
brew install im-select
```

4. Build:

```zsh
go build -o kana_to_eisu
```

## Usage

1. Go to System Preferences > Security & Privacy > Privacy > Accessibility
2. Check the box for Terminal (or your terminal application)
3. Run the program:

```zsh
./kana_to_eisu
```

## How it Works

- Uses `CGEventTap` to monitor keyboard input at the system level
- Detects key presses and modifier key (Shift, Ctrl, etc.) changes
- Automatically switches to English input mode after 3 seconds of keyboard inactivity
- Only switches when the current input mode is Japanese

## Notes

- Accessibility permissions are required
- Press Ctrl+C to exit the program

## License

MIT License
