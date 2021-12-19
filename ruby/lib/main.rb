require 'ui'
require 'loader'

class Main
  include UI
  include Loader

  def run
    query = prompt_user()
    data = load_data(query)
    results = search(data, query)
    output(results)
  end

  def search(data, query)
    data.filter { |record| record.match?(query) }
  end
end

