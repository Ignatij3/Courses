#include<iostream>
#include<cstring>
#include<algorithm>
#include<cmath>

#define ll long long
using namespace std;

const int MAX = 200;

    ll dp[MAX][8][8][8];

    char ans[MAX];
    const ll maxQ = 100000000000000LL;

ll go( int pos, int maxsum, int minsum, int tsum )
{
    if( maxsum > 7 ) return 0;
    if( minsum < 0 ) return 0;
    if( abs(tsum - maxsum) > 3 ) return 0;
    if( abs(tsum - minsum) > 3 ) return 0;

    if( pos == 0 ) return 1;
    
    if( dp[pos][maxsum][minsum][tsum] != -1 ) return dp[pos][maxsum][minsum][tsum];


    ll res = 0;

    res += go( pos - 1, maxsum, min( minsum, tsum - 1 ), tsum - 1 );
    res += go( pos - 1, max( maxsum, tsum + 1 ), minsum, tsum + 1 );

    if( res > maxQ ) res = maxQ;
    dp[pos][maxsum][minsum][tsum] = res;

return res;   
}

int main()
{
    memset( dp, -1, sizeof(dp));
    
    ll pos; scanf("%lld", &pos);

    int maxsum, minsum, tsum;

    int i;

    i = 1;

    while(1)
    {
     if( go( i-1, 3, 2, 2 ) + go( i-1, 4, 3, 4 )  >= pos )
     {
      break;
     }
     else pos -= go( i-1, 3, 2, 2 ) + go( i-1, 4, 3, 4 );

     ++i;
    }
     
    maxsum = 3; minsum = 3; tsum = 3;
   
    for( int j = i; j >= 1; --j )
    {
     if( go( j-1, maxsum, min(minsum, tsum - 1), tsum - 1 ) < pos )
     {
       pos -= go( j-1, maxsum, min(minsum, tsum- 1), tsum - 1 );
       ans[i-j] = 'b';
       tsum += 1; maxsum = max(maxsum, tsum);
     }  
     else {ans[i-j] = 'a'; tsum -= 1; minsum = min( minsum, tsum ); }
    }

    for( int j = 0; j < i; ++j ) printf("%c", ans[j]); printf("\n");
return 0;
}

