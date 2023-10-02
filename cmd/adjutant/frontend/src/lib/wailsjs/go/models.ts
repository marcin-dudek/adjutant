export namespace main {
	
	export class Author {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Author(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class Cover {
	    small: string;
	    medium: string;
	    large: string;
	
	    static createFrom(source: any = {}) {
	        return new Cover(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.small = source["small"];
	        this.medium = source["medium"];
	        this.large = source["large"];
	    }
	}
	export class Publisher {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Publisher(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class Book {
	    url: string;
	    key: string;
	    title: string;
	    authors: Author[];
	    publishers: Publisher[];
	    publish_date: string;
	    cover: Cover;
	
	    static createFrom(source: any = {}) {
	        return new Book(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.key = source["key"];
	        this.title = source["title"];
	        this.authors = this.convertValues(source["authors"], Author);
	        this.publishers = this.convertValues(source["publishers"], Publisher);
	        this.publish_date = source["publish_date"];
	        this.cover = this.convertValues(source["cover"], Cover);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	

}

