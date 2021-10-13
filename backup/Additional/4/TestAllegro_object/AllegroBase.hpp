#pragma once

#include <allegro5/allegro.h>
#include <allegro5/allegro_primitives.h>

class AllegroBase
{
    public:
        AllegroBase();
        virtual ~AllegroBase();

        bool Init( int screenWidth, int screenHeight, int fps );
        void Destroy();
        void Run();
        void Exit();

        virtual void Fps() = 0;
        virtual void Draw() = 0;

    protected:
        ALLEGRO_DISPLAY *alDisplay_;
        ALLEGRO_EVENT_QUEUE *alEventQueue_;
        ALLEGRO_TIMER *alTimer_;

    private:
        bool exit_;

};
