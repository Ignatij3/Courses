{
Written by Fyodor Menshikov 13.01.2004
11:38-12:03
array of coin sums removed at 12:29
I/O filename changed 21.04.2004
}
var
   pile:array[1..180]of integer;
   curgets:array[1..180,1..80]of longint;
   n,k,i:integer;
   curpile,notgreater,getpiles,pileafter:integer;
   canget,max,sum:longint;
begin
   assign(input,'coingame.in');
   reset(input);
   assign(output,'coingame.out');
   rewrite(output);
   read(n);
   for i:=1 to n do
      read(pile[i]);
   readln(k);
   for curpile:=n downto 1 do begin
      for notgreater:=k downto 1 do begin
         if notgreater>=n-curpile+1 then begin
            {gets all}
            sum:=0;
            for i:=curpile to n do
               sum:=sum+pile[i];
            curgets[curpile,notgreater]:=sum;
         end
         else begin
            max:=0;
            for getpiles:=1 to notgreater do begin
               pileafter:=curpile+getpiles;
               sum:=0;
               for i:=curpile to n do
                  sum:=sum+pile[i];
               canget:=sum-curgets[pileafter,getpiles];
               if canget>max then
                  max:=canget;
            end;
            curgets[curpile,notgreater]:=max;
         end;
      end;
   end;
   max:=0;
   for i:=1 to k do
      if curgets[1,i]>max then
         max:=curgets[1,i];
   writeln(max);
end.