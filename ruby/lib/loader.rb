require "json"

module Loader
  def load_data(query)
    rows = JSON.load(File.open("data/#{query.dataset}.json"))

    rows.map do |row|
      case query.dataset
      when "users"
        User.new(row)
      when "tickets"
        Ticket.new(row)
      else
        raise "Invalid dataset"
      end
    end
  end
end
