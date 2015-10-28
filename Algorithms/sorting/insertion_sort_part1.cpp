#include <map>
#include <set>
#include <list>
#include <cmath>
#include <ctime>
#include <deque>
#include <queue>
#include <stack>
#include <bitset>
#include <cstdio>
#include <vector>
#include <cstdlib>
#include <numeric>
#include <sstream>
#include <iostream>
#include <algorithm>

using namespace std;

void static print_arr(vector<int> arr){
    for(vector<int>::iterator it = arr.begin();
        it != arr.end();
        it++){
        
        cout << *it << " ";
    }
    cout << endl;
}

void static insertion_sort(vector<int>  arr) {
    int v = 0;
    int pos = 0;
    
    // finding the value needs to be sorted.
    for(int i=1; i<arr.size(); i++){
        if(arr[i] < arr[i-1]){
            v = arr[i];
            pos = i;
            break;
        }
    }
    
    // finding the position that V should be placed.
    while((v < arr[pos-1]) && (pos > 0)){
        arr[pos] = arr[pos-1];
        print_arr(arr);
        pos--;
    }
    
    // sorting done.
    arr[pos] = v;
    print_arr(arr);
}

int main(void) {
    vector <int>  _ar;
    int _ar_size;
    
    cin >> _ar_size;
    
    for(int _ar_i=0; _ar_i<_ar_size; _ar_i++) {
        int _ar_tmp;
        cin >> _ar_tmp;
        _ar.push_back(_ar_tmp); 
    }

    insertion_sort(_ar);
   
    return 0;
}

