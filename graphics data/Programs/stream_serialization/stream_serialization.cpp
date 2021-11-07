#include <iostream>
#include <map>
#include <string>
#include <vector>

class Serialization {
  protected:
    enum token_value {
        DELIM      = ',',
        ASSIGN     = '=',
        END        = ';',
        PARAM_LIST = ':'
    };
    std::string iter_string;

    char next()
    {
        if (iter_string.length() > 0)
        {
            char elem = iter_string[0];
            iter_string.erase(0, 1);
            return elem;
        }
        else
            return '\0';
    }

    std::string getValue()
    {
        std::string value;
        for (char token = next(); token != 0; token = next())
        {
            if (token == DELIM || token == END)
                break;
            value.push_back(token);
        }

        return value;
    }

    virtual void setValue(std::string member_name, std::string value) = 0;

  public:
    virtual std::string toString()                  = 0;
    virtual void fromString(const std::string& str) = 0;
};

class Square : public Serialization {
  private:
    double x, y;
    int side;

    void setValue(std::string member_name, std::string value)
    {
        if (member_name == "x")
        {
            x = atof(value.c_str());
        }
        else if (member_name == "y")
        {
            y = atof(value.c_str());
        }
        else if (member_name == "side")
        {
            side = atoi(value.c_str());
        }
    }

  public:
    Square(double xpos = 0, double ypos = 0, int sideLength = 1) :
        x(xpos), y(ypos), side(sideLength)
    {
        side = (side < 0) ? 0 : side;
    }

    std::string toString()
    {
        std::string repr;
        repr = "Square:x=" + std::to_string(x) + ",y=" + std::to_string(y) + ",side=" + std::to_string(side) + ";";
        return repr;
    }

    void fromString(const std::string& str)
    {
        iter_string = str;
        std::string buffer;

        for (char token = next(); token != 0 && token != END; token = next())
        {
            switch (token)
            {
            case ASSIGN:
                setValue(buffer, getValue());
                buffer = "";
                break;

            case PARAM_LIST:
                if (buffer != "Square")
                {
                    std::printf("Wrong object passed\n");
                    return;
                }

                buffer = "";
                break;

            default:
                buffer.push_back(token);
            }
        }
    }
};

class Circle {
  private:
    double x, y;
    int radius;

  public:
    Circle(double xpos = 0, double ypos = 0, int rad = 1);

    std::string toString()
    {
        std::string repr;
        repr = "Circle:x=" + std::to_string(x) + ",y=" + std::to_string(y) + ",radius=" + std::to_string(radius) + ";";
        return repr;
    }
};

int main()
{
    Square sqr(15, 15, 15);
    printf("before: %s\n", sqr.toString().c_str());
    sqr.fromString("Square:side=200,x=-13.333333,y=0;");
    printf("after: %s\n", sqr.toString().c_str());

    return 0;
}