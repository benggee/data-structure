public interface UF {
    int size();
    boolean unionContains(int p, int q);
    void union(int p, int q);
}