#include <iostream>
#include <string>

std::string extractFirstSubstring(std::string* str, const std::string& delim)
{
    std::size_t delim_pos = str->find(delim);
    if (delim_pos == -1)
        return "";

    std::string substring = str->substr(0, delim_pos + 1);
    str->erase(0, delim_pos + 1);

    return substring;
}

class Serialization {
  protected:
    std::string iter_string;
    enum token_value {
        DELIM  = ',',
        ASSIGN = '=',
        END    = ';',
    };

    Serialization() { }

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
        for (char token = next(); token != END; token = next())
        {
            if (token == DELIM)
                break;
            value.push_back(token);
        }

        return value;
    }
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
            if (token == ASSIGN)
            {
                setValue(buffer, getValue());
                buffer = "";
                continue;
            }
            buffer.push_back(token);
        }
    }
};

class Circle : public Serialization {
  private:
    double x, y;
    int radius;

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
        else if (member_name == "radius")
        {
            radius = atoi(value.c_str());
        }
    }

  public:
    Circle(double xpos = 0, double ypos = 0, int rad = 1) :
        x(xpos), y(ypos), radius(rad)
    {
        radius = (radius < 0) ? 0 : radius;
    }

    std::string toString()
    {
        std::string repr;
        repr = "Circle:x=" + std::to_string(x) + ",y=" + std::to_string(y) + ",radius=" + std::to_string(radius) + ";";
        return repr;
    }

    void fromString(const std::string& str)
    {
        iter_string = str;
        std::string buffer;

        for (char token = next(); token != 0 && token != END; token = next())
        {
            if (token == ASSIGN)
            {
                setValue(buffer, getValue());
                buffer = "";
                continue;
            }
            buffer.push_back(token);
        }
    }
};

class FigureFactory {
  private:
    std::string figureData;

  public:
    FigureFactory() { }

    FigureFactory(const std::string& str) :
        figureData(str) { }

    void addFigures(const std::string& str)
    {
        figureData += str;
    }

    void* getNext()
    {
        std::string curr_str = extractFirstSubstring(&figureData, ";");
        if (curr_str == "")
            return nullptr;
        std::string class_name = extractFirstSubstring(&curr_str, ":");

        if (class_name == "Square:")
        {
            Square* sqr = new Square;
            sqr->fromString(curr_str);
            return sqr;
        }
        else if (class_name == "Circle:")
        {
            Circle* circ = new Circle;
            circ->fromString(curr_str);
            return circ;
        }
        else
        {
            std::printf("Object cannot be parsed\n");
        }

        return nullptr;
    }
};

int main()
{
    FigureFactory fig("Square:side=200,x=-13.333333,y=0;Square:x=15,y=15,side=15;");
    Square* sqrptr = static_cast<Square*>(fig.getNext());
    if (sqrptr)
        printf("new: %s\n", sqrptr->toString().c_str());

    sqrptr = static_cast<Square*>(fig.getNext());
    if (sqrptr)
        printf("new: %s\n", sqrptr->toString().c_str());

    sqrptr = static_cast<Square*>(fig.getNext());
    if (sqrptr)
        printf("new: %s\n", sqrptr->toString().c_str());

    return 0;
}