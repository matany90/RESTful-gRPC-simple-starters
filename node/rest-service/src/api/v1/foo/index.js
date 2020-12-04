import { validateData } from "../../../middlewares/data-validation"

// Init Router
const router = require("express").Router({ mergeParams: true })

// GET foo
router.get("/", async (req, res) => {
  try {
    // Do whatever...

    res.json({ foo: "bar" })
  } catch (e) {
    console.error(e)
    res.status(500).json({ error: "unable to get foo." })
  }
})

// POST foo
router.post("/", validateData("foo"), async (req, res) => {
  try {
    // Do whatever...

    res.json({ bar: "zoo" })
  } catch (e) {
    console.error(e)
    res.status(500).json({ error: "unable to create foo." })
  }
})

// PUT foo
router.put("/", async (req, res) => {
  try {
    // Do whatever...
    res.json({ one: "two" })
  } catch (e) {
    console.error(e)
    res.status(500).json({ error: "unable to update foo." })
  }
})

// DELETE foo
router.delete("/", async (req, res) => {
  try {
    // Do whatever...

    res.json({ hoo: "boo" })
  } catch (e) {
    console.error(e)
    res.status(500).json({ error: "unable to delete foo." })
  }
})

// Export router
export default router
