{
Written by Fyodor Menshikov 13.01.2004
11:38-12:03
debug info written by 12:20
I/O filename changed 21.04.2004
}
{-$define debug}
var
   pile:array[1..180]of longint;
   coins:array[1..180]of longint;
   curgets:array[1..180,1..80]of longint;
   n,k,i:integer;
   curpile,notgreater,getpiles,pileafter:integer;
   canget,max:longint;
begin
   assign(input,'coingame.in');
   reset(input);
   assign(output,'coingame.out');
   rewrite(output);
   read(n);
   for i:=1 to n do
      read(pile[i]);
   readln(k);
   {$ifdef debug}
   writeln('n=',n);
   write('piles: ');
   for i:=1 to n do
      write(pile[i],' ');
   writeln;
   writeln('k=',k);
   {$endif}
   coins[n]:=pile[n];
   for i:=n-1 downto 1 do
      coins[i]:=coins[i+1]+pile[i];
   {$ifdef debug}
   writeln('coins in the pile and right ones:');
   for i:=1 to n do
      write(coins[i],' ');
   writeln;
   writeln;
   {$endif}
   for curpile:=n downto 1 do begin
      for notgreater:=k downto 1 do begin
         {$ifdef debug}
         write('first pile not got: ',curpile:3);
         writeln(' cannot get more than ',notgreater:2);
         {$endif}
         if notgreater>=n-curpile+1 then begin
            {gets all}
            {$ifdef debug}
            writeln(' getting all: ',coins[curpile]);
            {$endif}
            curgets[curpile,notgreater]:=coins[curpile];
         end
         else begin
            {$ifdef debug}
            writeln(' cannot get all.');
            {$endif}
            max:=0;
            for getpiles:=1 to notgreater do begin
               pileafter:=curpile+getpiles;
               canget:=coins[curpile]-curgets[pileafter,getpiles];
               {$ifdef debug}
               write(' if we get ',getpiles,' pile(s) ');
               write('opponent will get ');
               write(curgets[pileafter,getpiles],' coin(s)');
               write(' and we will get ');
               writeln(coins[curpile]-curgets[pileafter,getpiles],' coin(s)');
               {$endif}
               if canget>max then
                  max:=canget;
            end;
            {$ifdef debug}
            writeln(' max=',max);
            {$endif}
            curgets[curpile,notgreater]:=max;
         end;
      end;
   end;
   {$ifdef debug}
   writeln;
   write('Answer: ');
   {$endif}
   max:=0;
   for i:=1 to k do
      if curgets[1,i]>max then
         max:=curgets[1,i];
   writeln(max);
end.