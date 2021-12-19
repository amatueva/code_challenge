require 'record'

class User
  include Record

  FIELDS = ['_id', 'name', 'created_at', 'verified']

  def initialize(attrs)
    @attrs = attrs
  end

  def show
    heading = heading("User", @attrs['name'])
    body = row_as_table(@attrs)
    heading + body
  end
end
