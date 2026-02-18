# fixe

Fix English grammar mistakes in text using Groq API.

## Usage

```bash
fixe "<text>"
fixe --mimic theo "<text>"
```

## Flags

- `--mimic` - Output in a specific speaking style (currently: `theo`)

## Example

```bash
$ fixe "she dont like apples but she eat them anyway"
She doesn't like apples, but she eats them anyway.

$ fixe --mimic theo "this framework is really good but it have some performance issues"
Listen, this framework is genuinely good—a banger, really—but it has some performance issues that make it a pain to ship in production.
```

## Setup

1. Set your Groq API key in `.env.local`:
   ```
   GROQ_API_KEY=your-api-key
   ```

2. Build:
   ```bash
   go build -o fixe .
   ```
