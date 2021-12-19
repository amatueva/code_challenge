require "ticket"
require "user"
require "json"

module Validations
  def valid_dataset?(dataset)
    ["tickets", "users"].include?(dataset)
  end

  def valid_field?(dataset, field)
    case dataset
    when "users"
      User::FIELDS.include?(field)
    when "tickets"
      Ticket::FIELDS.include?(field)
    else
      raise "Invalid dataset"
    end
  end

  def valid_value?(value)
    JSON.parse(value)
    true
  rescue JSON::ParserError
    false
  end
end
