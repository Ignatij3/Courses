#ifndef GRAPHICS_SHAPES_H_
#define GRAPHICS_SHAPES_H_

#include <allegro5/allegro.h>
#include <allegro5/allegro_primitives.h>
#include <vector>

// move
bool AlmostEqual(double a, double b, double epsilon);

namespace shape
{
    const double PI = 3.14159265;

    enum Movement : bool {
        DYNAMIC = true,
        STATIC  = false
    };

    class Point {
      public:
        long double x, y;

        Point() noexcept;
        Point(long double x, long double y) noexcept;

        long double XDiff(const Point& rhs) const noexcept;
        long double YDiff(const Point& rhs) const noexcept;
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
        Vector(long double x1, long double y1, long double x2, long double y2) noexcept;
        Vector(const Vector& rhs) noexcept;

        void setAngle() noexcept;
        double getAngle() const noexcept;
        long double Magnitude() const noexcept;
        long double Slope() const noexcept;
        bool LiesBetween(const Vector& outsideVector) const noexcept;
        void SetVectors(Point point, Point vector) noexcept;
        void SetVectors(long double x1, long double y1, long double x2, long double y2) noexcept;
        bool Cross(const Vector& lineb) const noexcept;
        Vector Normal(bool left) const noexcept;

        long double HighestX() const noexcept;
        long double LowestX() const noexcept;
        long double HighestY() const noexcept;
        long double LowestY() const noexcept;

        int operator^(Vector& rhs) const noexcept;
        int operator*(Vector& rhs) const noexcept;
        Vector& operator=(const Vector& rhs) noexcept;
        Vector& operator=(Vector&& rhs) noexcept;
        bool operator==(const Vector& rhs) const noexcept;
    };

    // in circle, check for collision with vector pointing in move direction
    class Shape { //брать в расчёт массу и потерю скорости
      protected:
        bool dynamic;
        std::vector<Vector*> sides;
        std::pair<double, double> direction;

        virtual long double LeftX() const        = 0;
        virtual long double UpperY() const       = 0;
        virtual long double RightX() const       = 0;
        virtual long double LowerY() const       = 0;
        virtual void SetSides()                  = 0;
        virtual void SetSidesSetAngle()          = 0;
        virtual std::vector<Vector> GetNormals() = 0;

      public:
        double angle;
        Point centre;
        ALLEGRO_COLOR color = al_map_rgb(0, 0, 0);

        virtual std::vector<Vector*> GetSides() const  = 0;
        virtual const int constexpr sideAmount() const = 0;
        virtual void Move()                            = 0;
        virtual void Draw() const                      = 0;

        Shape(Movement move, Point centreCoords, double width, double height, double alpha) noexcept;
        Shape() noexcept;

        void SetDirection() noexcept;
        void Reflect(double otherVectorAngle) noexcept;
        void MovementToggle() noexcept;
        void MovementToggle(Movement move) noexcept;
        std::pair<bool, std::pair<const Vector*, const Vector*>> CollideWith(Shape* other) noexcept;
    };

    class Rectangle : public Shape {
      protected:
        double width, height;
        const short SIDE_AMOUNT = 4;
        std::vector<Vector*> sides;

        long double LeftX() const noexcept override;
        long double UpperY() const noexcept override;
        long double RightX() const noexcept override;
        long double LowerY() const noexcept override;
        void SetSides() noexcept override; // clockwise convention
        void SetSidesSetAngle() noexcept override;

      public:
        Rectangle(Movement move, Point centreCoords, double width, double height, double alpha);
        Rectangle(Movement move, long double centreX, long double centreY, double width, double height, double alpha);
        Rectangle() noexcept;
        ~Rectangle();

        Rectangle static InitFromStdin();
        std::vector<Vector*> GetSides() const noexcept override;
        const int sideAmount() const noexcept override;
        void Move() noexcept override;
        void Draw() const noexcept override;
        std::vector<Vector> GetNormals() override;
        Rectangle& operator=(const Rectangle& rhs) noexcept;
    };

    class Square : public Rectangle {
      public:
        Square(Movement move, Point centreCoords, double side, double alpha) noexcept;
        Square(Movement move, long double centreX, long double centreY, double side, double alpha) noexcept;

        Square static InitFromStdin();
    };
}

#endif // !GRAPHICS_SHAPES_H_