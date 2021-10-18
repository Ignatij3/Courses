#ifndef GRAPHICS_SHAPES_H_
#define GRAPHICS_SHAPES_H_

#include <vector>

// move
bool AlmostEqual(double a, double b, double epsilon);

namespace shape
{
    const double PI = 3.14159265;

    class Point {
      public:
        double x, y;

        Point() noexcept;
        Point(double x, double y) noexcept;

        double XDiff(const Point& rhs) const noexcept;
        double YDiff(const Point& rhs) const noexcept;
        bool operator==(const Point& rhs) const noexcept;
    };

    // vector defined with starting point and vector going from that point
    class Vector {
      private:
        double angle;

      public:
        ALLEGRO_COLOR color;
        Point a;
        Point b;

        Vector();
        Vector(Point point, Point vector) noexcept;
        Vector(double x1, double y1, double x2, double y2) noexcept;
        Vector(const Vector& rhs) noexcept;

        void setAngle() noexcept;
        double getAngle() const noexcept;
        double Magnitude() const noexcept;
        double Slope() const noexcept;
        bool LiesBetween(const Vector& outsideVector) const noexcept;
        void SetVectors(Point point, Point vector) noexcept;
        void SetVectors(double x1, double y1, double x2, double y2) noexcept;
        bool Cross(const Vector& lineb) const noexcept;

        double HighestX() const noexcept;
        double LowestX() const noexcept;
        double HighestY() const noexcept;
        double LowestY() const noexcept;

        int operator^(Vector& rhs) const noexcept;
        Vector& operator=(const Vector& rhs) noexcept;
        Vector& operator=(Vector&& rhs) noexcept;
        bool operator==(const Vector& rhs) const noexcept;
    };

    // in circle, check for collision with vector pointing in move direction
    class Shape { //брать в расчёт массу и потерю скорости //stationary option
      private:
        std::pair<bool, const Vector> LiesOnLine(const std::vector<Vector>& sides, const Point& angle) const noexcept;
        std::pair<const Vector, const Vector> FindSidesToReflect(std::vector<Vector>& shapeSides, std::vector<Vector>& otherShapeSides, int sideIndex, int otherSideIndex) const noexcept;

      protected:
        double angle;
        std::vector<Vector> sides;
        Point centre;
        std::pair<double, double> direction;

        virtual double LeftX() const  = 0;
        virtual double UpperY() const = 0;
        virtual double RightX() const = 0;
        virtual double LowerY() const = 0;
        virtual void SetSides()       = 0;
        virtual void SetAnglesSides() = 0;

      public:
        ALLEGRO_COLOR color = al_map_rgb(0, 0, 0);

        virtual std::vector<Vector> GetSides() const = 0;
        virtual const int sideAmount() const         = 0;
        virtual void Move()                          = 0;
        virtual void Draw() const                    = 0;

        Shape(Point centreCoords, double width, double height, double alpha) noexcept;

        void SetDirection(double alpha) noexcept;
        void Reflect(double otherVectorAngle) noexcept;
        std::pair<bool, std::pair<const Vector, const Vector>> CollideWith(const Shape* other) const noexcept;
        std::pair<bool, const Vector> CollideWith(const Vector* other_vector) const noexcept;
    };

    class Rectangle : public Shape {
    };

    class Square : public Shape {
      private:
        std::vector<Vector> sides;
        double edge; // distance from square's centre to edge

        double LeftX() const noexcept override;
        double UpperY() const noexcept override;
        double RightX() const noexcept override;
        double LowerY() const noexcept override;
        void SetSides() noexcept override;
        void SetAnglesSides() noexcept override;

      public:
        Square(Point centreCoords, double side, double alpha) noexcept;
        Square(double centreX, double centreY, double side, double alpha) noexcept;

        std::vector<Vector> GetSides() const noexcept override;
        const int sideAmount() const noexcept override;
        Square static InitFromStdin();
        void Move() noexcept override;
        void Draw() const noexcept override;
    };
}

#endif // !GRAPHICS_SHAPES_H_