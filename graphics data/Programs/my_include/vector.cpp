#include "shapes.hpp"

#include <cmath>
#include <iostream>

namespace shape
{

    Vector::Vector() { }

    Vector::Vector(Point point, Point vector) noexcept :
        a(point),
        b(vector)
    {
        setAngle();
    }

    Vector::Vector(long double x1, long double y1, long double x2, long double y2) noexcept :
        Vector(Point(x1, y1), Point(x2, y2))
    {
        setAngle();
    }

    Vector::Vector(const Vector& rhs) noexcept
    {
        a     = rhs.a;
        b     = rhs.b;
        angle = rhs.angle;
        color = rhs.color;
    }

    long double Vector::Slope() const noexcept
    {
        return (AlmostEqual(a.XDiff(b), 0, 1)) ? 0 : a.YDiff(b) / a.XDiff(b);
    }

    bool Vector::LiesBetween(const Vector& outsideVector) const noexcept
    {
        long double lowY  = outsideVector.LowestY();
        long double highY = outsideVector.HighestY();
        long double lowX  = outsideVector.LowestX();
        long double highX = outsideVector.HighestX();

        // printf("(%Lf >= %Lf && %Lf >= %Lf) && (%Lf >= %Lf && %Lf >= %Lf)\n", highX, this->b.x, this->b.x, lowX, highY, this->b.y, this->b.y, lowY);
        bool secondX = (highX > this->b.x || AlmostEqual(highX, this->b.x, 0.1)) && (this->b.x > lowX || AlmostEqual(this->b.x, lowX, 0.1));
        bool secondY = (highY > this->b.y || AlmostEqual(highY, this->b.y, 0.1)) && (this->b.y > lowY || AlmostEqual(this->b.y, lowY, 0.1));
        return secondX && secondY; // firstX && firstY
    }

    void Vector::setAngle() noexcept
    {
        angle = atan2(a.YDiff(b), a.XDiff(b)) * (180 / PI);
        angle = (angle < 0) ? 360 + angle : angle;
    }

    void Vector::SetVectors(Point point, Point vector) noexcept
    {
        a.x = point.x;
        a.y = point.y;
        b.x = point.x + vector.x;
        b.y = point.y + vector.y;
    }

    void Vector::SetVectors(long double x1, long double y1, long double x2, long double y2) noexcept
    {
        a.x = x1;
        a.y = y1;
        b.x = x1 + x2;
        b.y = y1 + y2;
    }

    double Vector::getAngle() const noexcept
    {
        return angle;
    }

    bool Vector::Cross(const Vector& lineb) const noexcept
    {
        auto sign = [](int x) -> char { return ((x > 0) ? '+' : ((x < 0) ? '-' : '0')); };
        int cross1, cross2, cross3, cross4;

        Vector ac(a, lineb.a);
        Vector ad(a, lineb.b);
        cross1 = *this ^ ac;
        cross2 = *this ^ ad;

        if (sign(cross1) == sign(cross2))
            return false;

        Vector ca(lineb.a, a);
        Vector cb(lineb.a, b);
        cross3 = lineb ^ ca;
        cross4 = lineb ^ cb;

        if (cross1 == 0 && cross2 == 0 && cross3 == 0 && cross4 == 0) // lies on same line
            return true;

        return !((sign(cross3) == sign(cross4)) || cross1 == 0 || cross2 == 0);
    }

    long double Vector::Magnitude() const noexcept
    {
        long double x_diff_squared = a.XDiff(b) * a.XDiff(b);
        long double y_diff_squared = a.YDiff(b) * a.YDiff(b);

        return sqrt(x_diff_squared + y_diff_squared);
    }

    long double Vector::HighestX() const noexcept
    {
        return (a.x > b.x) ? a.x : b.x;
    }

    long double Vector::LowestX() const noexcept
    {
        return (a.x < b.x) ? a.x : b.x;
    }

    long double Vector::HighestY() const noexcept
    {
        return (a.y > b.y) ? a.y : b.y;
    }

    long double Vector::LowestY() const noexcept
    {
        return (a.y < b.y) ? a.y : b.y;
    }

    // calculates determinant
    int Vector::operator^(Vector& rhs) const noexcept
    {
        int x1 = b.XDiff(a);
        int y1 = b.YDiff(a);
        int x2 = rhs.b.XDiff(rhs.a);
        int y2 = rhs.b.YDiff(rhs.a);

        return x1 * y2 - x2 * y1;
    }

    Vector& Vector::operator=(const Vector& rhs) noexcept
    {
        a = rhs.a;
        b = rhs.b;

        angle = rhs.angle;
        color = rhs.color;

        return *this;
    }

    Vector& Vector::operator=(Vector&& rhs) noexcept
    {
        a     = rhs.a;
        b     = rhs.b;
        angle = rhs.angle;
        color = rhs.color;

        return *this;
    }

    bool Vector::operator==(const Vector& rhs) const noexcept
    {
        return (a == rhs.a && b == rhs.b);
    }
}