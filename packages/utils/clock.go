/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package utils


// ClockWrapper represents wrapper of clock
type ClockWrapper struct {
}

// Now returns current time
func (cw *ClockWrapper) Now() time.Time { return time.Now() }
