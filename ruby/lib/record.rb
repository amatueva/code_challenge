module Record
  def match?(query)
    @attrs.has_key?(query.field) &&
      @attrs[query.field] == query.value
  end

  def heading(type, name)
    "# #{type} - #{name}\n\n"
  end

  def row_as_table(row)
    widest_key_width = row.keys.map(&:length).max
    row.map do |key, value|
      "%#{widest_key_width}s: %s" % [key, value]
    end.join("\n")
  end
end
