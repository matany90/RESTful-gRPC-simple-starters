import { validateData } from "../../../middlewares/data-validation"

// Init Router
const router = require("express").Router({ mergeParams: true })

// GET foo
router.get("/", async (req, res) => {
  res.json({ foo: "bar" })
})

// POST foo
router.post("/", validateData("foo"), async (req, res) => {
  try {
    res.json({ bar: "zoo" })
  } catch (e) {
    // console.error(e)
    res.status(500).json({ error: "unable to create foo." })
  }
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
