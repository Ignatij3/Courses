#include "shapes.hpp"

namespace shape
{
    template <class T>
    Point<T>::Point() :
        x(0), y(0) { }

    template <class T>
    Point<T>::Point(T x, T y) :
        x(x), y(y) { }

    template <class T>
    T Point<T>::XDiff(const Point& rhs) const
    {
        return this->x - rhs.x;
    }

    template <class T>
    T Point<T>::YDiff(const Point& rhs) const
    {
        return this->y - rhs.y;
    }
}