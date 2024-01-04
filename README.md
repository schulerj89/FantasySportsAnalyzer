
# Fantasy Sports Analyzer

Fantasy Sports Analyzer is a command-line tool written in Go, designed to process and analyze DraftKings export files. It reads CSV files containing player data, such as position, name, salary, game information, and average points per game, and provides useful insights and statistics based on this data.

## Features

- Read DraftKings export files in CSV format.
- Analyze player data to provide various statistics (e.g., average points per game, player salary analysis).
- Command-line interface for easy interaction.

## Getting Started

### Prerequisites

- Go (version 1.xx or later) installed on your machine.

### Installation

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/schulerjs89/FantasySportsAnalyzer.git
    ```

2. Navigate to the cloned directory:

    ```bash
    cd FantasySportsAnalyzer
    ```

3. Build the project (optional):

    ```bash
    go build .
    ```

### Usage

Run the program by specifying the path to the DraftKings export file:

```bash
go run . -file=<path_to_csv_file>
```

Or, if you have built the project:

```bash
./FantasySportsAnalyzer -file=<path_to_csv_file>
```

Replace `<path_to_csv_file>` with the actual path to your DraftKings CSV file.

### Output
Shows the average points per team
```
Team: MIL, Average Points: 16.01
Team: DEN, Average Points: 16.36
Team: GSW, Average Points: 17.35
Team: SAS, Average Points: 15.62
```

## License

This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details.
