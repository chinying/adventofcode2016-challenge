import collection.mutable.Queue
import collection.mutable.ListBuffer
import collection.mutable.HashSet
import collection.mutable.HashMap

case class Coordinate(x : Int, y : Int) {
  def neighbours(xe : Int, ye : Int) : List[Coordinate] = {
    var ls = new ListBuffer[Coordinate]()
    if (x > 0) ls.append(new Coordinate(x-1, y))
    if (y > 0) ls.append(new Coordinate(x, y-1))
    if (x < xe-1) ls.append(new Coordinate(x+1, y))
    if (y < ye-1) ls.append(new Coordinate(x, y+1))
    return ls.toList
  }
}
object Day13 {

  val input = 1352
  //val input = 10
  val goal = new Coordinate(31, 39)

  def bfs(src : Coordinate, dest: Coordinate, grid : Array[Array[Boolean]]) {
    val height = grid.length
    val width = grid.head.length
    var visited = new HashSet[Coordinate]
    var q = new Queue[Coordinate]
    var prev = new HashMap[Coordinate, Coordinate]
    q += src
    while (!q.isEmpty) {
      //at neighbours phase must check for both wall and boundaries, visited
      val front = q.dequeue
      val neighbours = front.neighbours(width, height).filter(n => {
        !grid(n.y)(n.x) && !visited(n) // not wall + not visited
      }).foreach(n => {
        visited += n
        q.enqueue(n)
        prev += (n -> front)
      })
    }

    var current = goal
    var path = new ListBuffer[Coordinate]
    path += current
    while (prev.contains(current) && (current.x != src.x && current.y != src.y)) {
      current = prev(current)
      path += current
    }
    println(path.length)
  }

  def hammingWeight(x : Int) : Int = {
    var v = x
    v = v - ((v>>1) & 0x55555555)
    v = (v & 0x33333333) + ((v>>2) & 0x33333333)
    ((v + (v>>4) & 0xF0F0F0F) * 0x1010101) >> 24
  }

  def printGrid(grid : Array[Array[Boolean]]) {
    grid.foreach(g => {
      g.foreach(c => {
        if (c) print("#")
        else print(".")
      })
      println()
    })
  }

  def isWall(x : Int, y : Int) = (hammingWeight(x*x + 3*x + 2*x*y + y + y*y + input) & 1) == 1

  def main(args: Array[String]): Unit = {
    var width = 60
    var height = 60
    var grid = Array.ofDim[Array[Boolean]](height)
    for (j<-0 until height) {
      grid(j) = Array.ofDim[Boolean](width)
      for (i<-0 until width) grid(j)(i) = isWall(i, j)
    }
    bfs(new Coordinate(1,1), goal, grid)
  }
}
