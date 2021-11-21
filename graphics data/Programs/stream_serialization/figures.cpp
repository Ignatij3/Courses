#include "figures.hpp"

#include "factory.hpp"

Serialization::Serialization() { }

char Serialization::next()
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

std::string Serialization::getValue()
{
    std::string value;
    for (char token = next(); token != END; token = next())
    {
        if (token == DELIM || token == '\0')
            break;
        value.push_back(token);
    }

    return value;
}

Figure::Figure(double xpos, double ypos) :
    x(xpos), y(ypos) { }

void Square::setValue(std::string member_name, std::string value)
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

Square::Square(double xpos, double ypos, int sideLength) :
    Figure(xpos, ypos),
    side(sideLength)
{
    side = (side < 0) ? 0 : side;
}

std::string Square::toString() const
{
    return "Square:x=" + std::to_string(x) + ",y=" + std::to_string(y) + ",side=" + std::to_string(side) + ";";
}

std::string Square::toFileString() const
{
    return "Square " + std::to_string(x) + " " + std::to_string(y) + " " + std::to_string(side) + ";";
}

void Square::fromString(const std::string& str)
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

void Square::fromString(std::stringstream& str)
{
    std::string value;

    std::getline(str, value, ' ');
    setValue("x", value);

    std::getline(str, value, ' ');
    setValue("y", value);

    std::getline(str, value, ' ');
    setValue("side", value);
}

void Circle::setValue(std::string member_name, std::string value)
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

Circle::Circle(double xpos, double ypos, int rad) :
    Figure(xpos, ypos), radius(rad)
{
    radius = (radius < 0) ? 0 : radius;
}

std::string Circle::toString() const
{
    return "Circle:x=" + std::to_string(x) + ",y=" + std::to_string(y) + ",radius=" + std::to_string(radius) + ";";
}

std::string Circle::toFileString() const
{
    return "Circle " + std::to_string(x) + " " + std::to_string(y) + " " + std::to_string(radius) + ";";
}

void Circle::fromString(const std::string& str)
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

void Circle::fromString(std::stringstream& str)
{
    std::string value;

    std::getline(str, value, ' ');
    setValue("x", value);

    std::getline(str, value, ' ');
    setValue("y", value);

    std::getline(str, value, ' ');
    setValue("radius", value);
}