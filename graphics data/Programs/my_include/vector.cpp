#include "shapes.hpp"

#include <cmath>
#include <iostream>

namespace shape
{
    void Vector::setAngle()
    {
        angle = atan2(a.YDiff(b), a.XDiff(b)) * 180 / PI;
        angle = (angle < 0) ? 360 + angle : angle;
    }

    Vector::Vector()
    {
    }

    Vector::Vector(Point<double> point, Point<double> vector) :
        a(point), b(vector)
    {
        setAngle();
    }

    Vector::Vector(double x1, double y1, double x2, double y2) :
        Vector(Point<double>(x1, y1), Point<double>(x2, y2))
    {
        setAngle();
    }

    void Vector::SetCoordinates(Point<double> point, Point<double> vector)
    {
        a.x = point.x;
        a.y = point.y;
        b.x = vector.x;
        b.y = vector.y;
        setAngle();
    }

    void Vector::SetCoordinates(double x1, double y1, double x2, double y2)
    {
        a.x = x1;
        a.y = y1;
        b.x = x2;
        b.y = y2;
        setAngle();
    }

    double Vector::getAngle() const
    {
        return angle;
    }

    bool Vector::Cross(const Vector& lineb) const
    {
        //printf("x1, y1: (%f, %f)\nx2, y2: (%f, %f)\n%f degrees\n", lineb.a.x, lineb.a.y, lineb.b.x, lineb.b.y, lineb.angle);

        auto sign = [](int x) -> char { return (x >= 0 ? '+' : '-'); };

        Vector ac(a, lineb.a);
        Vector ad(a, lineb.b);
        int cross1 = *this ^ ac;
        int cross2 = *this ^ ad;

        if (sign(cross1) == sign(cross2) || cross1 == 0 || cross2 == 0)
            return false;

        Vector ca(lineb.a, a);
        Vector da(lineb.a, b);
        cross1 = lineb ^ ca;
        cross2 = lineb ^ da;

        return !(sign(cross1) == sign(cross2) || cross1 == 0 || cross2 == 0);
    }

    double Vector::Magnitude()
    {
        double x_diff_squared = a.XDiff(b) * a.XDiff(b);
        double y_diff_squared = a.YDiff(b) * a.YDiff(b);

        return sqrt(x_diff_squared + y_diff_squared);
    }

    int Vector::operator^(Vector& rhs) const
    {
        int x1 = b.XDiff(a);
        int y1 = b.XDiff(a);
        int x2 = rhs.b.XDiff(rhs.a);
        int y2 = rhs.b.YDiff(rhs.a);

        return x1 * y2 - x2 * y1;
    }
}