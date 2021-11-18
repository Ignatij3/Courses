#define _USE_MATH_DEFINES
#include <cmath>

class Figure {
  protected:
    double width, height;

  public:
    Figure(double width, double height);
    static double totalArea(double acc, const Figure* fig)
    {
        return acc += fig->area();
    }
    virtual double area() const { return 0; }
};

class Rectangle : public Figure {
  public:
    Rectangle(double width, double height);
    double area() const;
};

class Square : public Figure {
  public:
    Square(double side);
    double area() const;
};

class Circle : public Figure {
  public:
    Circle(double radius);
    double area() const;
};