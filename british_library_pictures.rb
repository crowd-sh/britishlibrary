require 'rubygems'
require 'open-uri'
require 'nokogiri'

# { "user": { "id": "7857987@N06", "nsid": "7857987@N06",
#     "username": { "_content": "British Library" } }, "stat": "ok" }

1.upto(10200) do |i|
  doc = Nokogiri::HTML(open("http://www.flickr.com/photos/britishlibrary/page#{i}"))

  doc.css(".pc_img").each do |img|
    puts "#{i}, #{img.attributes['data-defer-src']}"
  end
end
