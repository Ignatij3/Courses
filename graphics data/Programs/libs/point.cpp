#include "shapes.hpp"

#include <iostream>

namespace shape
{
    Point::Point() noexcept :
        x(0), y(0) { }

    Point::Point(long double x, long double y) noexcept :
        x(x), y(y) { }

    long double Point::XDiff(const Point& rhs) const noexcept
    {
        return this->x - rhs.x;
    }

    long double Point::YDiff(const Point& rhs) const noexcept
    {
        return this->y - rhs.y;
    }

    bool Point::operator==(const Point& rhs) const noexcept
    {
        return (AlmostEqual(x, rhs.x, 0.4) && AlmostEqual(y, rhs.y, 0.4));
    }
}