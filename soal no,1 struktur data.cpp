#include <iostream> // Mengimpor pustaka untuk input dan output

using namespace std; // Menggunakan namespace std untuk memudahkan penulisan

int main() {
    int a; // Mendeklarasikan variabel a bertipe integer

    cout << "Masukkan sebuah angka: "; // Menampilkan pesan untuk meminta input dari pengguna
    cin >> a; // Membaca input dari pengguna dan menyimpannya ke dalam variabel a

    // Memeriksa apakah angka a sama dengan 5
    if (a == 5) {
        cout << "Ini adalah angka lima" << endl; // Menampilkan pesan jika a adalah 5
    } else {
        cout << "Ini bukan angka lima" << endl; // Menampilkan pesan jika a bukan 5
    }

    return 0; // Mengembalikan 0 menandakan program selesai
}

