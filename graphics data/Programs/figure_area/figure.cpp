#include "figure.hpp"

Figure::Figure(double width, double height) :
    width(width < 0 ? 0 : width),
    height(height < 0 ? 0 : height) { }

Rectangle::Rectangle(double width, double height) :
    Figure(width, height) { }

double Rectangle::area() const
{
    return width * height;
}

Square::Square(double side) :
    Figure(side, side) { }

double Square::area() const
{
    return width * width;
}

Circle::Circle(double radius) :
    Figure(radius, radius) { }

double Circle::area() const
{
    return 2 * M_PI * width;
}