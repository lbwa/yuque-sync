import handleAsync from 'func-async'
import { setFailed } from '@actions/core'
import { main } from './action'
;(async () => {
  const [, exception] = await handleAsync(main())
  if (exception) {
    setFailed(`Action failed with error ${exception.message}`)
  }
})()
