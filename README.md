## JSON Parser in Go Programming

**Objective:** Implementing a JSON parser in Go that checks JSON syntax and structure along with providing error messages for invalid JSON syantax/files.

**Features:**
- Returns exit code 0 for valid JSON; exit code 1 with error message for invalid JSON.
- Error messages include line and column numbers.
- Handles JSON elements: strings, booleans, null values, numbers (including negative), and object keys.

