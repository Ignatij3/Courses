#ifndef GRAPHICS_WINDOW_H_
#define GRAPHICS_WINDOW_H_

#include "shapes.hpp"

#include <exception>
#include <string>

namespace screen
{
    extern shape::Vector walls[4];

    class Window {
      private:
        struct BadInit : public std::exception {
            BadInit(const std::string& msg);

            const char* what() const throw();

          protected:
            std::string message;
        };

        std::pair<std::vector<shape::Vector*>, std::vector<shape::Shape*>> objects;

      public:
        static int fps;
        static int window_width;
        static int window_height;

        Window(int fps_, int win_width, int win_height);

        void Run(FpsCallback fpsCallback, DrawCallback drawCallback);
        void AddObject(shape::Vector*);
        void AddObject(shape::Shape* figure);
        void SetColor(shape::Shape* figure, ALLEGRO_COLOR col);
        void AddObject(shape::Shape* figure, ALLEGRO_COLOR col);
        void ManageCollisions() const;
        void MoveAll();
        void DrawAll() const;
    };
}

#endif // !GRAPHICS_WINDOW_H_