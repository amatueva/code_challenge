require 'record'

class Ticket
  include Record

  FIELDS = ['_id', 'created_at', 'type', 'subject', 'assignee_id', 'tags']

  def initialize(attrs)
    @attrs = attrs
  end

  def show
    heading = heading("Ticket", @attrs['subject'])
    body = row_as_table(@attrs)
    heading + body
  end
end
