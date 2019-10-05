# RUV m3u proxy
RUV provides a live stream on their website that uses a custom player that handles region blocking and such.
Initially it sends a GET request to an HTTP endpoint that returns the m3u8 playlist URL. This endpoint is region blocked.
The m3u8 playlist URL returned includes a UUID for the session. If you stop the session and try playing the same playlist it will not work (just show a static image).

This proxy handles getting a fresh session on every play. This is useful for watching RUV in any player by providing a static, known URL. **It does not circumvent any region blocking.**

## What it does
It fetches the m3u8 URL from RUV's api and returns a 302 redirect to that URL on every request. It just gets a session URL and hands it to the player.

