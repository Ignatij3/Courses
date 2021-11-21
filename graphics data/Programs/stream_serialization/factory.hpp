#ifndef SERIALIZATION_FACTORY_HPP
#include <sstream>
#include <string>

class FigureFactory {
  private:
    std::stringstream figureData;

  public:
    FigureFactory();
    FigureFactory(const std::string& str);
    void addFigures(const std::string& str);
    void* getNext();
    void* getNextStream(std::stringstream str);
};

#endif // SERIALIZATION_FACTORY_HPP