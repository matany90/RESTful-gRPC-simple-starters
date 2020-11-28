// Init Router
const router = require("express").Router({ mergeParams: true })

// GET foo
router.get("/", async (req, res) => {
  res.json({ foo: "bar" })
})

// POST foo
router.post("/", async (req, res) => {
  res.json({ bar: "zoo"})
})

// PUT foo
router.put("/", async (req, res) => {
  res.json({ one: "two" })
})

// DELETE foo
router.delete("/", async (req, res) => {
  res.json({ hoo: "boo" })
})

// Export router
export default router
