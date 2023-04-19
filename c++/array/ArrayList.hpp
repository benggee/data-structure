#define InitSize 10 // default size
#define IncSize 5

using namespace std;

template <typename T>
class ArrayList
{
public:
    ArrayList(int len = InitSize) {
        m_data = new T[len];
        m_length = 0;
        m_maxsize = len;
    }
    ~ArrayList() {
        delete[] m_data;
        m_length = 0; // set the length to 0
    }

public:
    // insert e at i
    bool Insert(int i, const T& e) {
        if (m_length >= m_maxsize)
        {
            return false; // the list is full
        }

        if (i <1 || i > (m_length + 1))
        {
            return false; // i is invalid
        }

        for (int j = m_length; j >= i; --j) {
            m_data[j] = m_data[j - 1];
        }

        m_data[i-1] = e;

        m_length++;

        return true;
    }

    // delete e at i
    bool Delete(int i) {
        if (m_length < 1) {
            return false; // the list is empty
        }

        if (i < 1 || i > m_length) {
            return false; // i is invalid
        }

        for (int j = i; j < m_length; ++j) {
            m_data[j-1] = m_data[j];
        }

        m_length--;

        return true;
    }

    // get the value of i
    bool GetElem(int i, T& e) {
        if (m_length < 1)
        {
            return false;
        }

        if (i < 1 || i > m_length) {
            return false;
        }

        e = m_data[i-1];

        return true;
    }

    // return the index of e
    int LocateElem(const T &e) {
        for (int i = 0; i < m_length; i++)
        {
            if (m_data[i] == e) {
                return i + 1;
            }
        }

        return -1;
    }

    // display the list
    void Display() {
        for (int i = 0; i < m_length; i++) {
            cout << m_data[i] << " ";
        }

        cout << endl;
    }


    // return the length of the list
    int Len() {
        return m_length;
    }

    // reverse the list
    void Reverse() {
        if (m_length <= 1)
        {
            return;
        }
        T tmp;
        for (int i = 0; i < m_length; ++i) {
            tmp = m_data[i];
            m_data[i] = m_data[m_length - i - 1];
            m_data[m_length - i - 1] = tmp;
        }
    }

private :
    // increase the size of the list
    void IncreaseSize() {
        T* p = m_data;
        m_data = new T[m_maxsize + IncSize];
        for (int i = 0; i < m_length; i++)
        {
            m_data[i] = p[i];
        }
        m_maxsize = m_maxsize + IncSize;
        delete[] p;
    }

private:
    T* m_data; // store the data
    int m_length; // the length of the list
    int m_maxsize; // the max size of the list
};