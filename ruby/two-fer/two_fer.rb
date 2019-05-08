# frozen_string_literal: true

# TwoFer returns one for you, one for me.
module TwoFer
  def self.two_fer(name = 'you')
    "One for #{name}, one for me."
  end
end
