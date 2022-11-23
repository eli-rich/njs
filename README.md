# njs.icu
A redirection url to many common things. Mainly focused on documentation.

- [njs.icu](#njsicu)
  - [When would I use this?](#when-would-i-use-this)
  - [How it works](#how-it-works)
  - [Current Sites](#current-sites)
  - [Add a site](#add-a-site)


## When would I use this?
Let's say your buddy is asking a really dumb JavaScript question about how to use `Array.map`, instead of doing the searching yourself you could just link him to: [njs.icu/mdn/array-map](https://njs.icu/mdn/array-map).

Or maybe your friend asks a really good question about what exactly Alan Turing did; just link him to [njs.icu/wiki/alan-turing](https://njs.icu/wiki/alan-turing).

The full list of supported websites are listed at [njs.icu](https://njs.icu). But I will also try to maintain them here.

## How it works
This project is heavily inspired by [mdn.io](insert_link). It uses duckduckgo to return the most relevant result.

Essentially, clicking on [njs.icu/mdn/object-assign](https://njs.icu/mdn/object-assign) is the same as:
1. Going to [duckduckgo](https://duckduckgo.com)
2. Searching for: "object-assign site:developer.mozilla.org"
3. Clicking on the first result.


This is great, but has the following side effects:

**Links may *change* over time.**

(Though it's worth noting that the link *should* always be relevant [according to duckduckgo])

**Some sites work better than others.**


For example, mdn is perfect for this because everything has its own page. So I could link something like [njs.icu/mdn/array-foreach](https://njs.icu/mdn/array-foreach).

In contrast, the NodeJS docs aren't super great with this project. The NodeJS docs have only a page for the module, so [njs.icu/node/fs-readfile](https://njs.icu/node/fs-readfile) and [njs.icu/node/fs-writefile](https://njs.icu/node/fs-writefile) will link to the same page.

## Current Sites
You can always view an up to date list of all the supported sites at [njs.icu](https://njs.icu). This page is generated on each config update.

- **/djs/** — [Discord.JS docs](https://discord.js.org/)
- **/express/** — [Express docs](https://expressjs.com/)
- **/fastify/** — [Fastify docs](https://fastify.io/)
- **/github/** — [Github repos](https://github.com)
- **/mdn/** — [Mozilla Developer Network docs](https://developer.mozilla.org/)
- **/mongo/** — [MongoDB docs](https://mongodb.com)
- **/mongoose/** — [Mongoose docs](https://mongoosejs.com/)
- **/next/** — [NextJS docs](https://nextjs.org)
- **/node** — [NodeJS docs](https://nodejs.org)
- **/npm/** — [NPM packages](https://npmjs.com)
- **/react/** — [React docs](https://reactjs.org)
- **/solid/** — [SolidJS docs](https://solidjs.com)
- **/svelte/** — [Svelte docs](https://svelte.dev)
- **/wiki/** — [Wikipedia](https://wikipedia.org)

## Add a site
Okay! Submit a pull request. Make sure the site fits these requirements:
1. Is somehow useful for developers.
2. Is for something that is widely used.
3. Does not use a misleading path.

Also, please make sure to update the readme and the `config.json` with the sites you want to add.

Your PR *will* be rejected if you rename or remove another site **without an excellent reason**.