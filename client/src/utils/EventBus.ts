/* eslint-disable @typescript-eslint/no-explicit-any */
export class EventBus {

	private events: Record<string, Array<(arg?: any) => void>> = {};

	/**
	 * Subscribes to an event types. fn will be called when the event is
	 * published. This method returns a function that can be invoked to
	 * unsubscribe.
	 */
	public subscribe(eventName: string, fn: (arg?: any) => void) {
		this.events[eventName] = this.events[eventName] || [];
		this.events[eventName].push(fn);
		return () => this.unsubscribe(eventName, fn);
	}

	private unsubscribe(eventName: string, fn: (arg?: any) => void) {
		if (this.events[eventName]) {
			for (let i = 0; i < this.events[eventName].length; i++) {
				if (this.events[eventName][i] === fn) {
					this.events[eventName].splice(i, 1);
					break;
				}
			}
		}
	}

	public publish(eventName: string, data: any) {
		if (this.events[eventName]) {
			this.events[eventName].forEach(fn => fn(data));
		}
	}
}

const mainBus = new EventBus();
export default mainBus;
