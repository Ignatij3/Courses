#ifndef SERIALIZATION_FIGURES_HPP
#include <sstream>
#include <string>

class Serialization {
  protected:
    std::string iter_string;
    enum token_value {
        DELIM  = ',',
        ASSIGN = '=',
        END    = ';',
    };

    Serialization();
    char next();
    std::string getValue();
};

class Figure : public Serialization {
  protected:
    double x, y;

    virtual void setValue(std::string member_name, std::string value) = 0;

  public:
    Figure(double xpos, double ypos);
    virtual std::string toString() const            = 0;
    virtual std::string toFileString() const        = 0;
    virtual void fromString(const std::string& str) = 0;
    virtual void fromString(std::stringstream& str) = 0;
};

class Square : public Figure {
  private:
    int side;

    void setValue(std::string member_name, std::string value);

  public:
    Square(double xpos = 0, double ypos = 0, int sideLength = 1);
    std::string toString() const;
    std::string toFileString() const;
    void fromString(const std::string& str);
    void fromString(std::stringstream& str);
};

class Circle : public Figure {
  private:
    int radius;

    void setValue(std::string member_name, std::string value);

  public:
    Circle(double xpos = 0, double ypos = 0, int rad = 1);
    std::string toString() const;
    std::string toFileString() const;
    void fromString(const std::string& str);
    void fromString(std::stringstream& str);
};

#endif // SERIALIZATION_FIGURES_HPP