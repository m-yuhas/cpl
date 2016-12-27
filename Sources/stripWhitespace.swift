// Strip Whitespace Except in Quotes
// (C) 2016 Michael Yuhas

func stripWhitespace( input_string: String ) -> String {
  var output_string = ""
  var inQuotes = false
  var prevChar = Character(" ")
  for char in input_string.characters {
    if char == " " || char == "\t" {
      prevChar = char
      if inQuotes {
        output_string.append(char)
      } else {
        continue
      }
    } else if char == "\"" || char == "“" || char == "”" {
      output_string.append(char)
      if inQuotes && prevChar != "\\" {
        prevChar = char
        inQuotes = false
        continue
      } else if !inQuotes {
        prevChar = char
        inQuotes = true
        continue
      } else {
        prevChar = char
        continue
      }
    } else {
      output_string.append(char)
      prevChar = char
    }
  }
  return output_string
}
