object NecklaceMatching {

    // a = ald | b = lda (start = 2, length=3)
    // 2,3,0 => 2+0-3=-1(<0) | 2+0=2 => b(2) = a
    // 2,3,1 => 2+1-3=0 | 2+1=3(>=3) => b(0) = l
    // 2,3,2 => 2+2-3=1 | 2+2=4(>=3) => b(1) = d
    def calcNewB(start: Int, length: Int, idx: Int): Int = {
        val sumAll = start + idx - length
        if (sumAll >= 0)
            return sumAll
        return start + idx
    }

    def main(args: Array[String]): Unit = {
        val a = "ald"
        val b = "lda"

        val bLength = b.length()

        val bStartIdxs = b.zipWithIndex.filter{ case (e, i) => e == a.charAt(0) }.map{ case (e, i) => i }

        val matchs = bStartIdxs.map{ startIdx =>
            (0 to bLength -1).map(i => calcNewB(startIdx, bLength, i)).map(i => b(i)).mkString("")
        }.count(newB => newB == a)

        if (matchs > 0)
            println("True")
        else
            println("False")
    }
}