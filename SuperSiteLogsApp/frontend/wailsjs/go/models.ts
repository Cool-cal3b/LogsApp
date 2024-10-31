export namespace handlers {
	
	export class CountResult {
	    errors: number;
	    messages: number;
	    success: boolean;
	    errorMessage: string;
	
	    static createFrom(source: any = {}) {
	        return new CountResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.errors = source["errors"];
	        this.messages = source["messages"];
	        this.success = source["success"];
	        this.errorMessage = source["errorMessage"];
	    }
	}
	export class LogViewOperationResult {
	
	
	    static createFrom(source: any = {}) {
	        return new LogViewOperationResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class LogsOperationResult {
	
	
	    static createFrom(source: any = {}) {
	        return new LogsOperationResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class OperationResult {
	
	
	    static createFrom(source: any = {}) {
	        return new OperationResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

