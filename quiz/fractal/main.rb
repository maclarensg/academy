def width(n)
    total = 0
    (0..n).each { |i|
        total += 2**i
        i+=1
    }
    total
end

class Array
    def print_pretty
        self.each do |row|
            row.each do |column|
                print column
            end
            puts
        end
    end
end

def height(n)
    2**n
end

def fractal(n)
    
    array =[]

    (height n).times{
        array << Array.new((width n), ".")
    }
    
    array.print_pretty
end


def main 
    if ARGV.size < 1 
        puts "Need a number"
        exit 1
    end
    fractal ARGV[0].to_i
end

if __FILE__ == $0
    main
end

