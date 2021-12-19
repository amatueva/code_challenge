require "validations"
require "query"
require "json"

module UI
  def prompt_user
    dataset = ask_for_dataset()
    field = ask_for_field(dataset)
    value = ask_for_value()
    Query.new(dataset, field, value)
  end

  def output(results)
    puts results.map(&:show).join("\n\n")
    puts
    puts "#{results.length} results."
  end

  private

  include Validations

  def ask_for_dataset()
    puts "Are you searching for tickets or users?"

    until(valid_dataset?(dataset = get_answer)) do
      puts "Please enter either tickets or users."
    end

    dataset
  end

  def ask_for_field(dataset)
    puts "Which field are you searching?"

    until(valid_field?(dataset, field = get_answer)) do
      puts "Please enter a valid field."
    end

    field
  end

  def ask_for_value()
    puts "What value are you searching for?"

    until(valid_value?(value = get_answer)) do
      puts "Please enter a valid value."
    end

    JSON.parse(value)
  end

  def get_answer
    answer = gets
    exit if answer.nil?
    answer.chomp
  end
end
